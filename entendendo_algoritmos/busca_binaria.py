from typing import Any


def busca_binaria(lista: Any, item: Any) -> int | None:
    baixo = 0
    alto = len(lista) - 1
    while baixo <= alto:
        meio = int((baixo + alto) / 2)
        chute = lista[meio]
        if chute == item:
            return meio
        if chute > item:
            alto = meio - 1
        else:
            baixo = meio + 1
    return None
