function novoCardTelaPrincipal(lista, propFoto, propNome, propCategoria, propEndereco, propNumero, propCidade, propEstado) {
    const novoCard = document.createElement('div')
    novoCard.classList.add('property-card')
    lista.append(novoCard)

    const container = document.createElement('figure')
    container.classList.add('property-info-cont')
    novoCard.append(container)


    const foto = document.createElement('img')
    foto.classList.add('property-foto')
    foto.src = `/img/propriedades/${propFoto}`
    container.append(foto)


    const propertyInfo = document.createElement('div')
    propertyInfo.classList.add('property-info')
    container.append(propertyInfo)

    const categoria = document.createElement('h5')
    categoria.classList.add('property-category')
    categoria.textContent = propCategoria
    propertyInfo.append(categoria)

    const nome = document.createElement('h3')
    nome.classList.add('property-name')
    nome.textContent = propNome
    propertyInfo.append(nome)

    const loc = document.createElement('p')
    loc.classList.add('property-location')
    loc.textContent = `${propCidade}, ${propEstado}`
    propertyInfo.append(loc)
}

document.addEventListener('DOMContentLoaded', async (e) => {
    e.preventDefault()

    const req = await fetch("/api/v1/propriedades/iniciais")
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
            novoCardTelaPrincipal(propertyList, prop.foto, prop.nome, prop.categoria,prop.endereco, prop.numero, prop.cidade, prop.estado, prop.descricao)
        })
    } else {
        document.getElementById('lista-propriedades-destaque').innerText = "Houve um erro em mostrar as propriedades em destaque"
        return
    }


    const elementos = document.querySelectorAll('.property-card');

    elementos.forEach(elemento => {
        elemento.addEventListener('click', (e) => {

            // Primeiro seleciona o card, depois pega o nome dele selecionando o elemento de nome que está dentro dele por classe
            const clicado = e.currentTarget;
            const nomeProp = clicado.querySelector('.property-name').textContent;
            // Como o nome está dentro de um h3
            alert(nomeProp)
        });
    });
})