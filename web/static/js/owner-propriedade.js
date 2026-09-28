document.addEventListener('DOMContentLoaded', async () => {
    // 1. Extrai o nome da propriedade da URL
    const pathParts = window.location.pathname.split('/');
    const nomeDaPropriedade = pathParts[pathParts.length - 1];

    const btnAdicionarQuarto = document.getElementById('btn-adicionar-quarto');
        btnAdicionarQuarto.addEventListener('click', () => {
            const urlAtual = window.location.pathname.replace(/\/$/, '');
            window.location.href = `${urlAtual}/adicionar`;
        });


    try {
        const response = await fetch(`/api/v1/propriedades/${nomeDaPropriedade}`);

        if (!response.ok) {
            throw new Error('Falha ao buscar os dados da propriedade.');
        }

        const prop = await response.json();

        document.getElementById('prop-nome').textContent = prop.nome;
        document.getElementById('prop-endereco').textContent = `📍 ${prop.endereco}, ${prop.numero} - ${prop.cidade}, ${prop.estado}`;

        document.getElementById('prop-categoria').textContent = prop.categoria;

        document.getElementById('prop-descricao').textContent = prop.descricao;
        document.getElementById('prop-avaliacao').textContent = `⭐ Avaliação: ${prop.avaliacao.toFixed(1)}`;

        if (prop.foto != null && prop.foto !== "") {
            document.getElementById('prop-foto').src = `/img/propriedades/${prop.foto}`;
        }

        if (prop.pet_friendly) {
            document.getElementById('prop-pet').style.display = 'inline-block';
        }

        const listaQuartos = document.getElementById('lista-quartos');
        const emptyMsg = document.getElementById('empty-rooms-msg');

        if (!prop.quartos || prop.quartos.length === 0) {
            // Mostra o texto "essa propriedade ainda não possui quartos"
            emptyMsg.style.display = 'block';
        } else {
            // Oculta a mensagem e renderiza a lista
            emptyMsg.style.display = 'none';

            prop.quartos.forEach(quarto => {
                const roomCard = document.createElement('div');
                roomCard.classList.add('property-card');

                roomCard.innerHTML = `
                    <div class="property-info">
                        <h3 class="property-name">${quarto.nome || 'Quarto Padrão'}</h3>
                        <p class="property-location">💰 R$ ${quarto.preco || '0.00'} / noite</p>
                        <p class="property-desc">${quarto.descricao || 'Sem descrição cadastrada.'}</p>
                    </div>
                    <div class="property-actions">
                        <button class="btn btn-outline btn-sm">Editar</button>
                        <button class="btn btn-danger btn-sm">Excluir</button>
                    </div>
                `;

                listaQuartos.appendChild(roomCard);
            });
        }

    } catch (erro) {
        console.error("Erro na comunicação com a API:", erro);
        document.getElementById('prop-nome').textContent = "Erro ao carregar os dados";
        document.getElementById('prop-descricao').textContent = "Verifique o backend e o console.";
    }
});