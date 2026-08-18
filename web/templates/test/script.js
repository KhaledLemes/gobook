document.addEventListener('DOMContentLoaded', () => {
    // --- Lógica do Dropdown de Hóspedes ---
    const guestsToggle = document.getElementById('guestsToggle');
    const guestsPopup = document.getElementById('guestsPopup');
    const guestsContainer = document.getElementById('guestsContainer');

    // Abre e fecha o popup
    guestsToggle.addEventListener('click', () => {
        guestsPopup.classList.toggle('active');
    });

    // Fecha ao clicar em outro lugar da tela
    document.addEventListener('click', (e) => {
        if (!guestsContainer.contains(e.target)) {
            guestsPopup.classList.remove('active');
        }
    });

    // --- Lógica de Validação de Datas ---
    const checkinInput = document.getElementById('checkin');
    const checkoutInput = document.getElementById('checkout');

    // Impede datas passadas no check-in
    const today = new Date().toISOString().split('T')[0];
    checkinInput.setAttribute('min', today);

    // Quando o usuário muda o check-in, o check-out mínimo é ajustado automaticamente
    checkinInput.addEventListener('change', () => {
        const checkinDate = new Date(checkinInput.value);

        // Adiciona 1 dia à data de check-in (impossibilitando checkout no mesmo dia)
        const nextDay = new Date(checkinDate);
        nextDay.setDate(checkinDate.getDate() + 1);

        const minCheckout = nextDay.toISOString().split('T')[0];
        checkoutInput.setAttribute('min', minCheckout);

        // Se o usuário já tinha colocado um checkout inválido, limpa e força o mínimo
        if (checkoutInput.value && checkoutInput.value <= checkinInput.value) {
            checkoutInput.value = minCheckout;
        }
    });
});

// --- Lógica dos Contadores de Hóspedes ---
let guests = {
    adult: 2,
    child: 0
};

window.updateGuest = function(type, change) {
    const newVal = guests[type] + change;

    // Regras de negócio: mínimo 1 adulto e mínimo 0 crianças
    if (type === 'adult' && newVal < 1) return;
    if (type === 'child' && newVal < 0) return;

    guests[type] = newVal;

    // Atualiza o texto no popup e no input condensado
    document.getElementById(`${type}Val`).innerText = guests[type];
    document.getElementById(`${type}CountDisplay`).innerText = guests[type];
};