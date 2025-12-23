import os
import subprocess
import sys

try:
    from dotenv import load_dotenv
except ImportError:
    print("dotenv not found. Installing...")
    subprocess.check_call([sys.executable, "-m", "pip", "install", "psycopg2-binary"])
    from dotenv import load_dotenv

load_dotenv()

try:
    import psycopg2
except ImportError:
    print("psycopg2 not found. Installing...")
    subprocess.check_call([sys.executable, "-m", "pip", "install", "psycopg2-binary"])
    import psycopg2

try:
    import requests
except ImportError:
    print("requests not found. Installing...")
    subprocess.check_call([sys.executable, "-m", "pip", "install", "requests"])
    import requests

DB_CONFIG = {
    "host": "localhost",
    "port": "5433",
    "dbname": "shumi",
    "user": "postgres",
    "password": os.getenv("POSTGRESQL_PASSWORD")
}

OLLAMA_API_URL = os.getenv("OLLAMA_API_URL")
OLLAMA_API_KEY = os.getenv("OLLAMA_API_KEY")
OLLAMA_MODEL = os.getenv("OLLAMA_MODEL")

def get_db_connection():
    return psycopg2.connect(**DB_CONFIG)

def load_tags_from_file(filepath='interests_tags.txt'):
    """Loads tags from a file, deduplicates them, and converts to lowercase."""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            tags = {line.strip() for line in f if line.strip()}
        return list(tags)
    except FileNotFoundError:
        print(f"Error: The file '{filepath}' was not found.")
        return []

def clear_interests_table(conn):
    """Deletes all records from the interests table."""
    with conn.cursor() as cur:
        print("Deleting all existing data from 'interests' table...")
        cur.execute("TRUNCATE TABLE interests RESTART IDENTITY CASCADE;")
        print("Table 'interests' cleared.")

def generate_embeddings_batch(tags):
    """Generates embeddings for a batch of tags."""
    try:
        response = requests.post(OLLAMA_API_URL, headers={'Authorization': f'Bearer {OLLAMA_API_KEY}'},
                                 json={"model": OLLAMA_MODEL, "input": [tag.lower() for tag in tags]})
        response.raise_for_status()
        return response.json().get("embeddings")
    except requests.exceptions.RequestException as e:
        print(f"Error generating embeddings: {e}")
        return None

def insert_embedding(conn, tag, embedding):
    """Inserts a tag and its embedding into the database."""
    with conn.cursor() as cur:
        embedding_str = str(embedding)
        cur.execute("INSERT INTO interests (tag, embedding) VALUES (%s, %s)", (tag, embedding_str))

def main():
    try:
        conn = get_db_connection()

        # Load tags from file
        tags_to_process = load_tags_from_file()
        
        if not tags_to_process:
            print("No tags found in 'interests_tags.txt' or file not found.")
            conn.close()
            return

        print(f"Found {len(tags_to_process)} unique tags to process.")

        print("Generating new embeddings for tags...")
        # The model expects a list of strings. The tags are already lowercased and unique.
        embeddings = generate_embeddings_batch(tags_to_process)

        if embeddings and len(embeddings) == len(tags_to_process):
            clear_interests_table(conn)
            print("Inserting new tags and embeddings into the database...")
            for i, embedding in enumerate(embeddings):
                tag = tags_to_process[i]
                insert_embedding(conn, tag, embedding)
            
            conn.commit()
            print(f"Successfully inserted {len(tags_to_process)} tags with embeddings.")
        else:
            print("Could not generate embeddings or there was a mismatch in the number of embeddings received.")

        conn.close()

    except psycopg2.Error as e:
        print(f"Database error: {e}")
    except Exception as e:
        print(f"An unexpected error occurred: {e}")


if __name__ == "__main__":
    main()
