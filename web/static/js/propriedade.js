document.addEventListener('DOMContentLoaded', async () => {
    // 1. Extrai o nome da propriedade direto da URL
    // Exemplo: se a URL for "http://localhost:8080/propriedade/Abraao-Marcos", ele pega "Abraao-Marcos"
    const pathParts = window.location.pathname.split('/');
    const nomeDaPropriedade = pathParts[pathParts.length - 1];

    try {
        // 2. Faz o fetch GET para a rota dinâmica do seu backend em Go
        const response = await fetch(`/api/v1/propriedades/${nomeDaPropriedade}`);

        if (!response.ok) {
            throw new Error('Falha ao buscar a propriedade');
        }

        const prop = await response.json();

        // 3. Preenche os campos estáticos do cabeçalho
        document.getElementById('title').textContent = prop.nome;
        document.getElementById('description').textContent = prop.descricao;
        document.getElementById('addr').textContent = `📍 ${prop.endereco}, ${prop.numero} - ${prop.cidade}, ${prop.estado}`;

        // Verifica se a categoria é um objeto aninhado na struct (ex: prop.categoria.nome) ou string
        document.getElementById('category').textContent = prop.categoria.nome || "Propriedade";

        document.getElementById('rating').textContent = prop.avaliacao.toFixed(1);

        // Preenche a foto principal
        if (prop.foto) {
            document.getElementById('main-photo').src = `/img/propriedades/${prop.foto}`;
        }

        // Validação Pet Friendly booleana da struct
        if (prop.pet_friendly) {
            const amenitiesBar = document.getElementById('amenities-bar');
            const petTag = document.createElement('span');
            petTag.classList.add('amenity-item');
            petTag.textContent = "🐾 Aceita pets";
            amenitiesBar.appendChild(petTag);
        }

        // 4. Preenche a tabela de quartos usando o array prop.quartos
        const tbody = document.getElementById('rooms-tbody');
        tbody.innerHTML = ""; // Limpa qualquer conteúdo residual

        if (prop.quartos && prop.quartos.length > 0) {
            prop.quartos.forEach(quarto => {
                // Cria a linha da tabela concatenando HTML dinâmico
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td class="room-detail-cell">
                        <h3 class="room-title">${quarto.nome || 'Quarto'}</h3>
                        <p class="room-description">${quarto.descricao || ''}</p>
                    </td>
                    <td class="text-center">👤👤</td>
                    <td class="price-cell">
                        <span class="price-amount">R$ ${quarto.preco || '0.00'}</span>
                    </td>
                    <td class="choices-cell">
                        <ul class="choices-list">
                            <li class="choice-success">✔️ Cancelamento grátis</li>
                        </ul>
                    </td>
                    <td class="action-cell">
                        <select class="select-room-count">
                            <option value="0">0</option>
                            <option value="1">1</option>
                        </select>
                        <button type="button" class="btn btn-primary btn-block mt-2">Vou reservar</button>
                    </td>
                `;
                tbody.appendChild(tr);
            });
        } else {
            // Se a propriedade ainda não tiver quartos cadastrados no banco
            tbody.innerHTML = `<tr><td colspan="5" style="text-align: center; padding: 2rem;">Nenhum quarto disponível no momento.</td></tr>`;
        }

    } catch (erro) {
        console.error("Erro no fetch da propriedade:", erro);
        document.getElementById('title').textContent = "Propriedade não encontrada";
        document.getElementById('description').textContent = "Verifique se o nome na URL está correto e se o servidor Go está rodando.";
    }
});