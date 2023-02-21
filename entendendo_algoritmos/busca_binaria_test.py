from busca_binaria import busca_binaria


def test_deve_retornar_o_indice_do_valor_caso_ele_exista_na_lista():
    numeros = [1, 2, 3, 4, 5]
    valor = 3

    resultado = busca_binaria(lista=numeros, item=valor)

    assert resultado == 2


def test_deve_retornar_None_caso_ele_nao_exista_na_lista():
    numeros = [1, 2, 3, 4, 5]
    valor = 8

    resultado = busca_binaria(lista=numeros, item=valor)

    assert resultado is None
