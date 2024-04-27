import openai
import os
import re
import ast
import re
import time

def quitar_comentarios_python(content):
    # Eliminar comentarios de una sola línea
    content = re.sub(r'#[^\n]*', '', content)
    # Eliminar comentarios de varias líneas
    content = re.sub(r'"""[\s\S]*?"""', '', content)
    content = re.sub(r"'''[\s\S]*?'''", '', content)
    return content