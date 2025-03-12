# Cache Simulator

**Cache Simulator** é uma aplicação desenvolvida em **Go (Golang)**, projetada para simular diversas políticas de cache, como **FIFO (First In, First Out)**, **LRU (Least Recently Used)**, e um **Cache de Múltiplos Níveis** (L1 e L2). A aplicação permite monitorar e exibir métricas de desempenho em tempo real, como **acertos** (hits) e **falhas** (misses), além de oferecer a possibilidade de expandir para uma interface web interativa.

## Funcionalidades

- **FIFO (First In, First Out)**: Implementação da política FIFO para o cache de nível 1 (RAM).
- **LRU (Least Recently Used)**: Implementação da política LRU para o cache de nível 1 (RAM).
- **Cache de Múltiplos Níveis**: Implementação de cache com dois níveis:
  - **L1 (RAM)**: Cache de alta velocidade para dados temporários.
  - **L2 (Disk)**: Cache persistido em disco, utilizado quando o cache de nível 1 (L1) atinge sua capacidade máxima.
- **Métricas de Desempenho**: Acompanhe o desempenho do cache com as métricas de acertos (hits) e falhas (misses).
- **Servidor Web (Em Desenvolvimento)**: Implementação de um servidor HTTP que fornece uma interface web com gráficos interativos e visualização em tempo real das métricas de cache.

## Como Usar

1. **Clone o repositório** para a sua máquina local:
   ```bash
   git clone https://github.com/evertonreis1/cache-simulator.git
   ```

2. **Navegue até o diretório do projeto**:
   ```bash
   cd cache-simulator
   ```

3. **Instale as dependências do Go** (se necessário):
   ```bash
   go mod tidy
   ```

4. **Execute o servidor Go**:
   ```bash
   go run main.go
   ```

## Tecnologias

- **Go (Golang)**: Linguagem de programação utilizada para implementar o sistema de cache e suas políticas.
- **Chart.js**: Biblioteca JavaScript para criar gráficos interativos na interface web (em desenvolvimento).

## Contribuições

Contribuições são **muito bem-vindas**! Se você tiver sugestões, melhorias ou correções, sinta-se à vontade para abrir um **Pull Request** ou **Issue**. Sua contribuição será altamente apreciada!
