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
choosedLanguages = ['python3', 'java', 'go', 'rust', 'php', 'javascript']
TEMPLATE_TABLE_TAG = "{%-- TABLE --%}"
TEMPLATE_COUNT_TAG = "{%-- COUNT --%}"
LANGUAGES = {
  'python3': 'py',
  'java': 'java',
  'go': 'go',
  'rust': 'rs',
  'php': 'php',
  'javascript': 'js'
}

def valid_filename(filename):
 if filename == '.git' or filename == '.github' or filename == 'readme-generator.py' or filename == 'README.md' or filename == 'README-TEMPLATE.md' :
   return False
 return True

def format_filename(filename, language):
  words = filename.split('-')
  index = words[0]
  title = ' '.join(words[1:])
  url = './%s/%s' % (filename , filename+'.'+LANGUAGES[language])
  logger.info("[%s] - %s - %s" % (index, title, url))
  return index, title, url

def generate_markdown_table(questions):
  logger.info('Starting generate markdown table.')
  table = """
  | ID   | Title | Difficulty | Java | Python | Golang | PHP | Rust | Javascript |
  | :----: | :----- | :----- | :----: | :------: | :------: | :------: | :------: | :------: |
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

  def to_markdown(self):
    java_url = "[java](%s)" % (self.solutions['java']) if 'java' in self.solutions and os.path.exists(self.solutions['java']) else "-"
    python_url = "[python3](%s)" % (self.solutions['python3']) if 'python3' in self.solutions and os.path.exists(self.solutions['python3']) else "-"
    go_url = "[go](%s)" % (self.solutions['go']) if 'go' in self.solutions and os.path.exists(self.solutions['go']) else "-"
    php_url = "[php](%s)" % (self.solutions['php']) if 'php' in self.solutions and os.path.exists(self.solutions['php']) else "-"
    rust_url = "[rust](%s)" % (self.solutions['rust']) if 'rust' in self.solutions and os.path.exists(self.solutions['rust']) else "-"
    javascript_url = "[javascript](%s)" % (self.solutions['javascript']) if 'javascript' in self.solutions and os.path.exists(self.solutions['javascript']) else "-"
    title_url = "[%s](%s/README.md)" % (self.title, self.title_url)

    difficulty = ''
    with open('./'+self.title_url+'/README.md') as f:
     readme_problem = f.readline()
     difficulty_reg = "<h3>(.*?)</h3>"
     problem_url = re.search(r'href=[\'"]?([^\'" >]+)', readme_problem)
     if problem_url:
      difficulty = "[%s](%s)" % (re.findall(difficulty_reg, readme_problem)[0], problem_url.group(1))
     else:
      difficulty = "%s" % (re.findall(difficulty_reg, readme_problem)[0])
    markdown = "|%s|%s|%s|%s|%s|%s|%s|%s|%s|" % (self.index, title_url, difficulty, java_url, python_url, go_url, php_url, rust_url, javascript_url)
    return markdown

if __name__ == "__main__":
  logger.info('Starting read local solutions and generate question list')
  
  # - get all file names
  questions = {}
  for language in choosedLanguages:
    
    folderName = "%s" % (os.path.abspath(os.path.dirname(__file__)))
    for filename in os.listdir(folderName):
     
     if valid_filename(filename):
      index, title, url = format_filename(filename, language)
      if index not in questions:
        questions[index] = Question(index, title, filename)
      questions[index].add_solution(language, url)
    sortedQuestions = collections.OrderedDict(sorted(questions.items()))
    logger.info('Finish generate question list, there are %d questions' % (len(sortedQuestions)))  
    
    # - Read template and generate README.md
    templatePath = "%s/%s" % (os.path.abspath(os.path.dirname(__file__)), 'README-template.md')
    readmePath = "%s/%s" % (os.path.abspath(os.path.dirname(__file__)), 'README.md')

    with open('README-template.md', 'r') as template, open('README.md', 'w') as f:
      logger.info('Starting write file')
      for line in template.readlines():
        if TEMPLATE_TABLE_TAG in line:
          table = generate_markdown_table(sortedQuestions) + "\n"
          f.write(table)
        elif TEMPLATE_COUNT_TAG in line:
          count = 'Problem totals: %s' % len(sortedQuestions)
          f.write(count)
        else:
          f.write(line)
    logger.info('Finished!')
