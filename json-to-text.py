import json, os

# script to convert scrapped json gh issues to text
# structure of converted text file
# [Issue Title]
# [Comment 1]
# [Comment n]

def extract_text(name):
    # Specify the path to your JSON file
    json_file = '/tmp/data/issues/' + name + ".json"

    # Read JSON data from the file
    with open(json_file, 'r') as file:
        json_data = file.read()

    # Load JSON data
    data = json.loads(json_data)

    # Extract title
    text = data["title"] + "\n"

    # Extract body text from the main post
    text = text + data["body"] + "\n"

    # Extract body text from comments
    comments = data["comments"]["edges"]
    for comment in comments:
        text =  text + comment["node"]["body"] + "\n"

    # Store body text in a text file
    text_file = 'data/issues/' + name + ".txt"
    with open(text_file, "w") as file:
        file.write(text)

# Specify the path to the folder containing JSON files
folder_path = '/tmp/data/issues'

# Get a list of all files in the folder
all_files = os.listdir(folder_path)

# Filter only JSON files
json_files = [file for file in all_files if file.endswith('.json')]

# Extract names without the ".json" extension
json_file_names = [os.path.splitext(file)[0] for file in json_files]

for n in json_file_names:
    extract_text(n)