#!/usr/bin/python
import os
import logging
import collections
import re

# - config logger
loggerFormat = '%(asctime)-15s %(message)s'
logging.basicConfig(format=loggerFormat)
logger = logging.getLogger('readme-generator')
logger.setLevel(logging.INFO)

# - Config
choosedLanguages = ['python3', 'java', 'go', 'rust', 'php', 'javascript', 'typescript', 'sql']
TEMPLATE_TABLE_TAG = "{%-- TABLE --%}"
TEMPLATE_COUNT_TAG = "{%-- COUNT --%}"
LANGUAGES = {
  'python3': 'py',
  'java': 'java',
  'go': 'go',
  'rust': 'rs',
  'php': 'php',
  'javascript': 'js',
  'typescript': 'ts',
  'sql': 'sql'
}

def valid_filename(filename):
    excluded_files = {'.git', '.github', 'readme-generator.py', 'README.md', 
                     'README-TEMPLATE.md', 'stats.json', '_config.yml'}
    return filename not in excluded_files

def format_filename(filename, language):
    words = filename.split('-')
    index = words[0]
    title = ' '.join(words[1:])
    url = os.path.join('.', filename, f"{filename}.{LANGUAGES[language]}")
    logger.info(f"[{index}] - {title} - {url}")
    return index, title, url

def generate_markdown_table(questions):
    logger.info('Starting generate markdown table.')
    table = """
  | ID   | Title | Difficulty | Java | Python3 | Golang | PHP | Rust | Javascript | Typescript | SQL |
  | :----: | :----- | :----- | :----: | :------: | :------: | :------: | :------: | :------: | :------: | :------: |
  """

    items = list(map(lambda item: item.to_markdown(), questions.values()))
    table = table + "\n".join(items)
    logger.info('Finish generate markdown table.')
    return table

class Question:
    def __init__(self, index, title, title_url):
        self.index = index
        self.title = title
        self.title_url = title_url
        self.solutions = {}

    def add_solution(self, language, url):
        self.solutions[language] = url

    def _get_solution_url(self, language):
        if language in self.solutions and os.path.exists(self.solutions[language]):
            return f"[{language}]({self.solutions[language]})"
        return "-"

    def _get_difficulty(self):
        readme_path = os.path.join('.', self.title_url, 'README.md')
        try:
            with open(readme_path) as f:
                readme_problem = f.readline()
                difficulty_reg = "<h3>(.*?)</h3>"
                problem_url = re.search(r'href=[\'"]?([^\'" >]+)', readme_problem)
                if problem_url:
                    return "[%s](%s)" % (re.findall(difficulty_reg, readme_problem)[0], problem_url.group(1))
                return "%s" % (re.findall(difficulty_reg, readme_problem)[0])
        except Exception as e:
            logger.error(f"Error reading difficulty from {readme_path}: {str(e)}")
            return "Unknown"

    def to_markdown(self):
        solution_urls = {lang: self._get_solution_url(lang) for lang in choosedLanguages}
        title_url = f"[{self.title}]({self.title_url}/README.md)"
        difficulty = self._get_difficulty()
        
        return "|{}|{}|{}|{}|{}|{}|{}|{}|{}|{}|{}|".format(
            self.index, title_url, difficulty,
            solution_urls['java'],
            solution_urls['python3'],
            solution_urls['go'],
            solution_urls['php'],
            solution_urls['rust'],
            solution_urls['javascript'],
            solution_urls['typescript'],
            solution_urls['sql']
        )

def main():
    try:
        logger.info('Starting read local solutions and generate question list')
        questions = {}
        folder_name = os.path.abspath(os.path.dirname(__file__))

        # Scan files once and process solutions
        for filename in os.listdir(folder_name):
            if valid_filename(filename):
                for language in choosedLanguages:
                    index, title, url = format_filename(filename, language)
                    if index not in questions:
                        questions[index] = Question(index, title, filename)
                    questions[index].add_solution(language, url)

        sorted_questions = collections.OrderedDict(sorted(questions.items()))
        logger.info(f'Found {len(sorted_questions)} questions')

        # Generate README
        template_path = os.path.join(folder_name, 'README-TEMPLATE.md')
        readme_path = os.path.join(folder_name, 'README.md')

        if not os.path.exists(template_path):
            raise FileNotFoundError(f"Template file not found: {template_path}")

        with open(template_path, 'r') as template, open(readme_path, 'w') as f:
            logger.info('Starting write file')
            for line in template.readlines():
                if TEMPLATE_TABLE_TAG in line:
                    table = generate_markdown_table(sorted_questions) + "\n"
                    f.write(table)
                elif TEMPLATE_COUNT_TAG in line:
                    f.write(f'Problem totals: {len(sorted_questions)}')
                else:
                    f.write(line)
        logger.info('Finished!')

    except Exception as e:
        logger.error(f"An error occurred: {str(e)}")
        exit(1)

if __name__ == "__main__":
    main()
