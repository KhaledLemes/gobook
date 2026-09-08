document.addEventListener('DOMContentLoaded', (e) => {
    e.preventDefault()

    const butCriar = document.getElementById('bt-enviar')
    butCriar.addEventListener('click', async (e) => {
        const nome = document.getElementById('nome')
        const descricao = document.getElementById('descricao')
        const categoria = document.getElementById('categoria')
        const endereco = document.getElementById('endereco')
        const numero = document.getElementById('numero')
        const cidade = document.getElementById('cidade')
        const estado = document.getElementById('estado')
        const petFriendly = document.getElementById('pet_friendly')
        const err = document.getElementById('err')

        const req = await fetch("/api/v1/propriedades", {
            method: 'POST',
            body: JSON.stringify({
                "nome": nome.value,
                "descricao": descricao.value,
                "categoria": categoria.value,
                "endereco": endereco.value,
                "numero": numero.value,
                "cidade": cidade.value,
                "estado": estado.value,
                "pet_friendly": petFriendly.checked,
            })
        })

        if (req.status !== 200) {
            const data = await req.json()
            err.style.color = 'red'
            err.innerText = ''
            err.innerText = data.error
        } else {
            window.location.replace('/propriedades/minhas')
        }
    })
});