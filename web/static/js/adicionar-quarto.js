document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('form-adicionar-quarto');
    const btnSalvar = document.getElementById('btn-salvar-quarto');
    const errorMessage = document.getElementById('err');

    form.addEventListener('submit', async function(e) {
        e.preventDefault();

        const propriedadeID = new URLSearchParams(window.location.search).get('p');
        if (!propriedadeID || !/^[1-9]\d*$/.test(propriedadeID)) {
            errorMessage.textContent = 'ID da propriedade inválido.';
            return;
        }

        const naoReembolsavel = document.getElementById('nao-reembolsavel').checked;
        const reembolso = document.getElementById('reembolso').value;
        const quarto = {
            nome: document.getElementById('nome').value,
            descricao: document.getElementById('descricao').value,
            valor_noite: Number(document.getElementById('valor_noite').value),
            disponivel: Number(document.getElementById('disponivel').value),
        };
        if (naoReembolsavel) {
            quarto.reembolso = '0001-01-01T00:00:00Z';
        } else if (reembolso) {
            quarto.reembolso = `${reembolso}T00:00:00Z`;
        }

        btnSalvar.disabled = true;
        errorMessage.textContent = '';

        try {
            const response = await fetch(`/api/v1/quartos?propID=${encodeURIComponent(propriedadeID)}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(quarto),
            });

            if (!response.ok) {
                const result = await response.json();
                errorMessage.textContent = result.error || 'Não foi possível salvar o quarto.';
                return;
            }

            const urlPropriedade = new URL(window.location.href);
            urlPropriedade.pathname = urlPropriedade.pathname.replace(/\/adicionar\/?$/, '');
            urlPropriedade.search = '';
            window.location.replace(urlPropriedade.toString());
        } catch (error) {
            console.error('Erro ao salvar o quarto:', error);
            errorMessage.textContent = 'Erro de comunicação ao salvar o quarto.';
        } finally {
            btnSalvar.disabled = false;
        }
    });
});
