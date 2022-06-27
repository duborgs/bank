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
  
  # INICIALIZAÇÃO COM MYSQL LOCAL
    . Copiar o arquivo script.sql em src/pkg/bd e executar no mysql local
    . alterar o arquivo db.go (src/pkg/bd) para db, err := sql.Open("mysql", "usuario:senha@tcp(localhost)/q2bank")
    . cd src/pkg/domains
    . go run router.go
    
  
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
    . Para usar banco local
        db, err := sql.Open("mysql", "user:password@tcp(localhost)/banco") - src/pkg/db/db.go
        
# OBSERVAÇÕES
    . Optei por não persistir os dados do container em um volume por conta do problema de autenticação do windows, somente administradores da máquina podem excluir a pasta. Linux não tem esse problema, pois pode utilizar o comando sudo chmod para resolver essa questãpo.
    . Fiz uma valiadação para que um usuário comum não faça transferência para a própria conta.
    . Código em inglês para todos os desenvolvedores (brasileiros e estrangeiros) entender, caso tenha necessidade, reutilizar o código.
    
# FUTURAS ATUALIZAÇÕES
    . Melhoria no versionamento das Branc (BR-01, BR-01.1);
    . Mudar o recebimento do endpoint transaction de r.FormValue para um Body(Json);
    . Documentação em inglês;
    . Separação de usuário e pessoa;
    . Transferência entre bancos;
    . Criptografia dos dados sensiveis no banco de dados;
    . Implementação do JWT.  
  
