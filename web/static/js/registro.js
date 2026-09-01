function getRole() {
    return document.querySelector('input[name="tipo_conta"]:checked').value
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
    const nome = document.getElementById('nome');
    const nomeMeio = document.getElementById('nome_meio');
    const nomeUltimo = document.getElementById('nome_ultimo');
    const nascimento = document.getElementById('nascimento');
    const tel = document.getElementById('tel');


    const stp1err = document.getElementById('stp1-err')
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
        console.log(roleSelecionado.value)

        e.preventDefault()
        const formatedDate = new Date(nascimento.value).toISOString()
        const req = await fetch("/api/v1/usuarios", {
            method: 'POST',
            body: JSON.stringify({
                "email": email.value,
                "senha": senha.value,
                "nome": nome.value,
                "nome_meio": nomeUltimo.value,
                "nome_ultimo": nomeMeio.value,
                "nascimento": formatedDate,
                "tel": tel.value,
                "role": getRole()
            })
        })
    });

});