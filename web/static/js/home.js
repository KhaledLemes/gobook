function novoCardTelaPrincipal(lista, propNome, propEndereco, propNumero, propCidade, propEstado, propDescricao) {
    const novoCard = document.createElement('div')
    novoCard.classList.add('property-card')
    lista.append(novoCard)

    const propertyInfo = document.createElement('div')
    propertyInfo.classList.add('property-info')
    novoCard.append(propertyInfo)

    const nome = document.createElement('h3')
    nome.classList.add('property-name')
    nome.textContent = propNome
    propertyInfo.append(nome)

    const loc = document.createElement('p')
    loc.classList.add('property-location')
    loc.textContent = `${propEndereco}, ${propNumero} - ${propCidade}, ${propEstado}`
    propertyInfo.append(loc)

    const desc = document.createElement('span')
    desc.classList.add('property-desc')
    desc.textContent = propDescricao
    propertyInfo.append(desc)


    const acoes = document.createElement('div')
    acoes.classList.add('property-actions')
    novoCard.append(acoes)

    const excluir = document.createElement('button')
    excluir.classList.add('btn')
    excluir.classList.add('btn-ver-mais')
    excluir.classList.add('btn-sm')
    excluir.textContent = "Ver mais"
    acoes.append(excluir)
}

document.addEventListener('DOMContentLoaded', async (e) => {
    e.preventDefault()

    const req = await fetch("/api/v1/propriedades")
    const body = await req.json()

    if (req.status === 200) {
        const data = JSON.parse(JSON.stringify(body))
        const propertyList = document.getElementById('lista-propriedades-destaque')
        if (data === null) {
            propertyList.style.textAlign = 'center'
            propertyList.style.color = '#666666'
            propertyList.innerText = "Houve um erro em mostrar as propriedades em destaque"
            return
        }
        data.forEach(prop => {
            novoCardTelaPrincipal(propertyList, prop.nome, prop.endereco, prop.numero, prop.cidade, prop.estado, prop.descricao)
        })
    } else {
        document.getElementById('lista-propriedades-destaque').innerText = "Houve um erro em mostrar as propriedades em destaque"
    }
})