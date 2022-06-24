# ENDPOINTS
  # TRANSACTIONS 
    .POST: /transaction
    .GET:/getTransactions
  # USERS:
    .GET: /getUserID7
    .POST: /createUser
  
# INICIALIZANDO APLICAÇÃO COM DOCKER
    . Na pasta raiz
    . docker-compose up -d
  
     . cd src/pkg/domains
     . go run router.go  
  
  # INICIALIZAÇÃO LOCAL
    . Copiar o arquivo script.sql em src/pkg/bd e executar no mysql local
    . src/pkg/bd db, err := sql.Open("mysql", "usuario:senha@tcp(localhost)/q2bank")
  
# BANCO DE DADOS
    . SGBD utilizado: MYSQL;
    . Query's em caixa baixa para ser reconhecido tanto no linux quanto no windows (linux não faz sensitive case);
    . Duas tabelas, uma users e outra transactions;

# CONTAINER MYSQL (DOCKER)
    . Usado para containeriar um serviço mysql;
    . Comandos para acessar o banco e visualizar registros e executar query's;
        "docker ps" para exibir os contaiers;
        "docker exec -it IDContainer bash" para acessar o terminal da base de dados;
        "mysql -u User -p Password".

        PRONTO! AGORA É SÓ REALIZAR QUERY'S!

        CTRL+D para sair do terminal
    .Para usar banco local
        db, err := sql.Open("mysql", "user:password@tcp(localhost)/banco") - src/pkg/db/db.go
