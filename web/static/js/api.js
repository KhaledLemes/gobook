document.addEventListener('DOMContentLoaded', async (e) => {
    e.preventDefault()

    const visitantes = document.getElementById('botoes-visitante')
    const logado = document.getElementById('area-logado')
    const nomeUser = document.getElementById('nome-usuario')

    const req = await fetch("/api/v1/me")
    const data = await req.json()

    if (req.status === 200) {
        visitantes.style.display = 'none'
        logado.style.display = 'flex'

        nomeUser.innerText = data.nome
    } else {
        visitantes.style.display = 'flex'
        logado.style.display = 'none'
    }

    if (data.error != null && data.error === "Sessão expirada, por favor, faça login novamente") {
        alert(data.error)
    }


    const sair = document.getElementById('logout')
    sair.addEventListener('click', async (e) => {
        const req = await fetch("/api/v1/logout")

        if (req.status === 200) {
            window.location.replace("/")
        }
    })

    const painelBtn = document.getElementById('painel')
    painelBtn.addEventListener('click', (e) => {
        switch (data.role) {
            case "owner":
                window.location.assign("/propriedades/minhas")
                break
            case "guest":
                window.location.assign("/reservas/minhas")
                break
            case "admin":
                window.location.assign("/admin/painel")
        }
        if (data.role === "owner") {
        }
    })
})

