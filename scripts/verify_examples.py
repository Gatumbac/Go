#!/usr/bin/env python3
"""Construye y ejecuta ejemplos principales; no recorre código del estudiante."""
from pathlib import Path
import argparse
import json
import os
import re
import subprocess
import tempfile
ROOT = Path(__file__).resolve().parents[1]
parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument('--race', action='store_true', help='Instrumentar ejemplos de concurrencia')
args = parser.parse_args()
manifest = json.loads((ROOT/'docs/lesson-manifest.json').read_text())
if args.race:
    manifest = [d for d in manifest if d['module'] == '05-concurrency']
with tempfile.TemporaryDirectory(prefix='go-course-check-') as tmp:
    env = dict(os.environ)
    env.setdefault('GOCACHE', '/tmp/go-from-zero-gocache')
    env.setdefault('GOTOOLCHAIN', 'local')
    env.setdefault('GOMODCACHE', '/tmp/go-from-zero-modcache')
    names = [Path(d['example']).name for d in manifest]
    if len(set(names)) != len(names):
        raise RuntimeError('Nombres de binarios duplicados')
    subprocess.run(['go', 'build', *(['-race'] if args.race else []), '-o', tmp + os.sep, *['./' + d['example'] for d in manifest]], cwd=ROOT, env=env, check=True, timeout=180)
    for i, d in enumerate(manifest):
        binary = str(Path(tmp)/Path(d['example']).name)
        result = subprocess.run([binary], env=env, capture_output=True, text=True, check=True, timeout=20)
        output = result.stdout.strip()
        expected = d.get('expected_stdout')
        if expected is not None:
            if output != expected:
                raise AssertionError(f'{d["example"]}: {output!r} != {expected!r}')
        elif d['slug'] == 'compilation':
            if not re.fullmatch(r'[a-z0-9]+/[a-z0-9]+', output):
                raise AssertionError('Destino de compilación inválido')
        elif d['slug'] == 'concurrency':
            if sorted(output.splitlines()) != ['alpha', 'beta']:
                raise AssertionError('Faltan resultados concurrentes')
        elif d['slug'] == 'gc':
            if not re.fullmatch(r'heap=\d+ gc-cycles=\d+', output):
                raise AssertionError('Métrica GC inválida')
        elif d['slug'] == 'profiling':
            if not re.fullmatch(r'allocations/run=\d+\.\d+', output):
                raise AssertionError('Métrica de asignaciones inválida')
        else:
            raise AssertionError('Ejemplo sin contrato de salida')
        print(f'OK {d["example"]}: {output!r}', flush=True)
print(f'OK: {len(manifest)} ejemplos construidos y terminados; race={args.race}.')
