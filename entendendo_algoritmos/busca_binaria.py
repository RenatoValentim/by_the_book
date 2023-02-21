from typing import Any, List


def busca_binaria(lista: List[Any], item: Any) -> int | None:
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
