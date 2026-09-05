document.addEventListener('DOMContentLoaded', async (e) => {
    e.preventDefault()

    const visitantes = document.getElementById('botoes-visitante')
    const logado = document.getElementById('area-logado')
    const nomeUser = document.getElementById('nome-usuario')

    const req = await fetch("/api/v1/me")
    if (req.status === 200) {
        visitantes.style.display = 'none'
        logado.style.display = 'flex'

        const data = await req.json()
        nomeUser.innerText = data.nome
    } else {
        visitantes.style.display = 'flex'
        logado.style.display = 'none'
    }


    const sair = document.getElementById('logout')
    sair.addEventListener('click', async (e) => {
        const req = await fetch("/api/v1/logout")

        if (req.status === 200) {
            window.location.replace("/")
        }
    })

})

