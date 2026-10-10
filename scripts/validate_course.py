#!/usr/bin/env python3
"""Valida documentación sin modificar implementaciones del estudiante."""
from pathlib import Path
import collections
import json
import re
import sys
import unicodedata
import xml.etree.ElementTree as ET
from urllib.parse import unquote, urlsplit
ROOT = Path(__file__).resolve().parents[1]
errors = []
def check(condition, message):
    if not condition:
        errors.append(message)
def anchors(text):
    seen = collections.Counter()
    out = set()
    for title in re.findall(r'^#{1,6}\s+(.+)$', text, re.M):
        title = re.sub(r'\[([^]]+)\]\([^)]*\)', r'\1', title)
        title = ''.join(c for c in title.lower() if c in ' -_' or unicodedata.category(c)[0] in 'LN')
        slug = title.replace(' ', '-')
        n = seen[slug]
        seen[slug] += 1
        out.add(slug if not n else f'{slug}-{n}')
    return out
manifest = json.loads((ROOT/'docs/lesson-manifest.json').read_text())
check(len(manifest) == 49, 'Deben existir 49 lecciones')
check(len({d['lesson'] for d in manifest}) == 49, 'Lecciones duplicadas')
check(sorted(collections.Counter(d['module'] for d in manifest).values()) == [7,7,8,9,9,9], 'Distribución incorrecta')
for d in manifest:
    p = ROOT/d['lesson']
    check(p.exists(), f'Falta {p}')
    if not p.exists():
        continue
    text = p.read_text()
    for section in d['sections']:
        match = re.search(r'^## '+re.escape(section)+r'\n(.*?)(?=^## |\Z)', text, re.M|re.S)
        check(bool(match and len(match.group(1).strip()) > 40), f'{p}: sección vacía {section}')
    check(len(re.findall(r'^\d+\. ', text, re.M)) in range(2,5), f'{p}: preguntas fuera de 2–4')
    code = re.search(r'```go\n(.*?)\n```', text, re.S)
    source = ROOT/d['example']/'main.go'
    check(source.exists(), f'Falta ejemplo de {p}')
    if source.exists():
        check(bool(code and code.group(1).strip() == source.read_text().strip()), f'Código Markdown difiere: {p}')
    check((p.parent/'self-assessment'/f'{p.stem}.md').exists(), f'Falta autoevaluación: {p}')
    for kind in ('guided','challenge'):
        exercise = p.parent/'exercises'/p.stem/kind
        for name in ('README.md','HINTS.md'):
            check((exercise/name).exists(), f'Falta {exercise/name}')
        for test in exercise.rglob('*_test.go'):
            check(test.read_text().startswith('//go:build exercise\n'), f'Falta tag: {test}')
        for source in exercise.rglob('*.go'):
            if not source.name.endswith('_test.go'):
                print(f'Source del estudiante presente, no modificado ni auditado: {source.relative_to(ROOT)}')
links = 0
for p in ROOT.rglob('*.md'):
    if '.git' in p.parts:
        continue
    text = p.read_text()
    check(not re.search(r'\b(TODO|TBD|contenido pendiente de generar)\b', text), f'Pendiente editorial: {p}')
    prose = re.sub(r'```.*?```', '', text, flags=re.S)
    prose = re.sub(r'`[^`\n]+`', '', prose)
    for target in re.findall(r'!?\[[^\]\n]*\]\(([^)\n]+)\)', prose):
        target = target.strip().strip('<>')
        parsed = urlsplit(target)
        if parsed.scheme or target.startswith('//'):
            continue
        links += 1
        resolved = (p.parent/unquote(parsed.path)).resolve() if parsed.path else p
        check(resolved.exists(), f'{p.relative_to(ROOT)}: link roto {target}')
        if resolved.exists() and resolved.suffix == '.md' and parsed.fragment:
            check(unquote(parsed.fragment) in anchors(resolved.read_text()), f'Ancla rota {p}: {target}')
for p in (ROOT/'.vscode').glob('*.json'):
    json.loads(p.read_text())
for p in (ROOT/'assets').glob('*.svg'):
    ET.parse(p)
    check('http://' not in p.read_text().replace('http://www.w3.org/2000/svg',''), f'Recurso remoto SVG: {p}')
if errors:
    print('\n'.join(errors), file=sys.stderr)
    sys.exit(1)
print(f'OK: 49 lecciones, {links} enlaces locales, secciones, ejemplos, enunciados, JSON y SVG')
print('Mermaid: no certificado por parser en este script; consultar docs/verification.md.')
