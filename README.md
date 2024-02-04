<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body>
    <h1>Aplicativo de Receitas em Go</h1>
    <p>Este é um projeto de aplicativo de receitas desenvolvido em Go. O aplicativo permite aos usuários visualizar, adicionar, editar e excluir receitas.</p>
    <h2>Funcionalidades</h2>
    <ul>
        <li>Listar todas as receitas disponíveis.</li>
        <li>Visualizar detalhes de uma receita específica.</li>
        <li>Adicionar uma nova receita.</li>
        <li>Editar uma receita existente.</li>
        <li>Excluir uma receita.</li>
    </ul>
    <h2>Tecnologias Utilizadas</h2>
    <ul>
        <li>Go (Golang)</li>
        <li>PostgreSQL</li>
        <li>Bibliotecas adicionais: <code>github.com/lib/pq</code> (driver PostgreSQL)</li>
    </ul>
    <h2>Estrutura do Projeto</h2>
    <ul>
        <li><strong>cmd/</strong>: Contém o arquivo <code>main.go</code> que inicia a aplicação.</li>
        <li><strong>internal/api/</strong>: Contém os manipuladores HTTP da API.</li>
        <li><strong>internal/database/</strong>: Contém a implementação da conexão com o banco de dados PostgreSQL.</li>
        <li><strong>internal/models/</strong>: Contém os modelos de dados da aplicação.</li>
        <li><strong>migrations/</strong>: Contém arquivos de migração para gerenciar o esquema do banco de dados.</li>
        <li><strong>config/</strong>: Contém as configurações da aplicação.</li>
    </ul>
    <h2>Instalação e Execução</h2>
    <ol>
        <li>Certifique-se de ter o Go e o PostgreSQL instalados em sua máquina.</li>
        <li>Clone este repositório: <code>git clone https://github.com/seu-usuario/aplicativo-receitas-go.git</code></li>
        <li>Navegue até o diretório do projeto: <code>cd aplicativo-receitas-go</code></li>
        <li>Execute as migrações do banco de dados: <code>go run ./migrations/*.go</code></li>
        <li>Configure as variáveis de ambiente necessárias, como a string de conexão do banco de dados.</li>
        <li>Inicie o servidor: <code>go run cmd/main.go</code></li>
        <li>Acesse o aplicativo em <code>http://localhost:8080</code>.</li>
    </ol>
    <h2>Licença</h2>
    <p>Este projeto é licenciado sob a <a href="LICENSE">MIT License</a>.</p>
</body>
</html>