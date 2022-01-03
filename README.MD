# Trainning Basic Go lang Architecture

## Architecture

 - Adapter
   - Ex:. User 
     - User / controller
       - port
       - Ex:. Grpc, Api, Kafka ...
     - User / presenter
     - User / repository
       - port


 - Adapter
   - Ex: controller(Api, Web), gateway(Repository) ...
     

 - Usecase
   - port
   - system
      - Ex:. add product
   - Interator
     - Ex:. user
   - dto
     - output
     - input
   - Repository
     - interface repository


 - Entities
   - models
     - Ex:. User & Product
   - port
     - repository

> each infrastructure with its controller

> port -> interface

> adapter -> implementations interface


**controllers** receive the external data

**presenter** -> transform in the right format "external output".

**adapter** -> translate between the outside world and business rules

**dto input** -> receives data from controllers "bind"

**dto output** -> intended format


```markdown
# infrastructure
data =! input

# usecase
output =! data external

```
Exemplo : .
* Designer Patterns & SOLID