let stp1err = document.getElementById('stp1-err')
let stp2err = document.getElementById('stp2-err')



function getValue(el, porId) {
    if (porId) {
        return document.getElementById(el).value
    }
    return document.querySelector(el).value
}

document.addEventListener('DOMContentLoaded', (e) => {
    e.preventDefault()
    const step1 = document.getElementById('step1');
    const step2 = document.getElementById('step2');

    const btnNext = document.getElementById('btnNext');
    const btnPrev = document.getElementById('btnPrev');
    const btnFinish = document.getElementById('finish');

    const email = document.getElementById('email');
    const senha = document.getElementById('senha');
    const nascimento = document.getElementById('nascimento');

    stp1err.style.color = "red"

    btnNext.addEventListener('click', (e) => {
        e.preventDefault()
        if (email.value === "" || senha.value === "") {
            stp1err.innerText = ""
            stp1err.innerText = "Campo de e-mail ou de senha estão vazios"
            return
        }
        if (senha.value.length < 8) {
            stp1err.innerText = ""
            stp1err.innerText = "Senha deve ter pelo menos 8 dígitos"
            return;
        }
        step1.classList.remove('active');
        step2.classList.add('active');
    });

    btnPrev.addEventListener('click', () => {
        step2.classList.remove('active');
        step1.classList.add('active');
    });


    btnFinish.addEventListener('click', async (e) => {

        e.preventDefault()
        const formatedDate = new Date(nascimento.value).toISOString()
        const req = await fetch("/api/v1/usuarios", {
            method: 'POST',
            body: JSON.stringify({
                "email": getValue('email', true),
                "senha": getValue('senha', true),
                "nome": getValue('nome', true),
                "nome_meio": getValue('nome_meio', true),
                "nome_ultimo": getValue('nome_ultimo', true),
                "nascimento": formatedDate,
                "tel": getValue('tel', true),
                "role": getValue('input[name="tipo_conta"]:checked', false)
            })
        })

        if (req.status !== 200) {
            const data = await req.json()
            stp2err.style.color = 'red'
            stp2err.innerText = data.error
        } else {
            const log = await fetch("/api/v1/login", {
                body: JSON.stringify({
                    "email": getValue('email', true),
                    "senha": getValue('senha', true)
                }),
                method: 'POST',
            })
            if (log.status === 200) {
                window.location.replace("/home")
            } else {
                stp2err.style.color = 'green'
                stp2err.innerText = "Você conseguiu se registrar mas houve uma falha ao fazer o seu login. Volte para a página inicial e tente fazer login com as credenciais cadastradas"
            }
        }
    });


});