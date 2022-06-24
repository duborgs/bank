# Banco de dados
    . SGBD utilizado: MYSQL;
    . Query's em caixa baixa para ser reconhecido tanto no linux quanto no windows (linux não faz sensitive case);
    . Duas tabelas, uma users e outra transactions;

# Docker
    . Usado para utilizar um serviço mysql, assim, não é necessário ter o mysql na máquina. Apenas o oDcker e GO;
    . Comandos para acessar o banco e visualizar registros e executar query's;
        "docker ps" para exibir os contaiers;
        "docker exec -it IDContainer bash" para acessar o terminal da base de dados;
        "mysql -u User -p Password".

        PRONTO! AGORA É SÓ REALIZAR QUERY'S!

        CTRL+D para sair do terminal
    .Para usar banco local
        db, err := sql.Open("mysql", "user:password@tcp(localhost)/banco") - src/pkg/db/db.go

# Volume - *data*
    . Configurado no docker-compose para persistir os dados e utilizar sempre que subir outros containers;
        Possui credenciais de usuário e registros da base de dados.