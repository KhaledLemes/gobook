function novoCard(lista, propNome, propCidade, propEstado, propDescricao) {
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
    loc.textContent = `${propCidade}, ${propEstado}`
    propertyInfo.append(loc)

    const desc = document.createElement('span')
    desc.classList.add('property-desc')
    desc.textContent = propDescricao
    propertyInfo.append(desc)

    const acoes = document.createElement('div')
    acoes.classList.add('property-actions')
    novoCard.append(acoes)

    const editar = document.createElement('button')
    editar.classList.add('btn')
    editar.classList.add('btn-outline')
    editar.classList.add('btn-sm')
    editar.textContent = "Editar"
    acoes.append(editar)

    const excluir = document.createElement('button')
    excluir.classList.add('btn')
    excluir.classList.add('btn-danger')
    excluir.classList.add('btn-sm')
    excluir.textContent = "Excluir"
    acoes.append(excluir)
}

document.addEventListener('DOMContentLoaded', async (e) => {
    const p = document.getElementById('teste')


    e.preventDefault()
    const req = await fetch("/api/v1/propriedades/minhas")
    const body = await req.json()

    if (req.status === 200) {
        const data = JSON.parse(JSON.stringify(body))
        const propertyList = document.getElementById('lista-propriedades')
        if (data === null) {
            propertyList.style.textAlign = 'center'
            propertyList.style.color = '#666666'
            propertyList.innerText = "Você não possui propriedades :("
            return
        }
        data.forEach(prop => {
            novoCard(propertyList, prop.nome, prop.cidade, prop.estado, prop.descricao)
        })
    } else {
        alert("erro")
    }
});

