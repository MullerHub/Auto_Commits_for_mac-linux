# 🌙 Luna — Gerador de Commits com IA

Luna gera mensagens de commit concisas para seus arquivos staged usando a API Google Gemini 2.0 Flash.

## ✨ Funcionalidades

- **Commits por arquivo**: um commit para cada arquivo staged
- **Gemini 2.0 Flash**: resumos gerados por IA a partir dos diffs
- **Prefixos convencionais**: adiciona prefixo se não existir
- **Controle de tamanho**: alvo < 60 chars, máximo configurável (padrão 72)
- **Filtros inteligentes**: ignora binários e imagens comuns
- **Emojis opcionais**: habilite com `-e`

## Como funciona

1. Coleta arquivos staged via `git diff --cached --name-only`
2. Envia cada diff para `gemini-2.0-flash:generateContent`
3. Se a resposta não tiver prefixo conhecido, escolhe aleatoriamente um de: `chore:`, `refactor:`, `feat:`, `fix:`, `docs:`, `test:`
4. Se `-e` estiver ativo, adiciona emoji aleatório
5. Trunca para `maxCommitLength` e faz commit com `git commit -m <message> -- <file>`

## Requisitos

- macOS | Linux | Sistemas UNIX-like
- Git instalado e disponível no `PATH`
- Chave da API Google Gemini (`https://aistudio.google.com/app/apikey`)

## Instalação

### Opção A — Usar binário pré-compilado (`bin/luna`)

1. Copie `bin/luna` para um diretório do sistema, por exemplo `/usr/local/bin`:

```bash
sudo cp bin/luna /usr/local/bin/
sudo chmod +x /usr/local/bin/luna
```

Ou adicione a pasta onde o binário está ao seu `PATH` (adicione a linha ao `~/.zshrc` ou `~/.bashrc`):

```bash
export PATH="$PATH:/caminho/para/sua/pasta"
source ~/.zshrc
```

### Opção B — Compilar do código-fonte (Go)

```bash
go build -o ./bin/luna main.go
```

Ou use o script helper:

```bash
./build.sh
```

## Configuração

Luna lê configuração de arquivos de projeto e global:

- Projeto: `.lunacfg` (na raiz do repositório ou pai mais próximo)
- Global: `.lunarc` (no diretório home do usuário)

Prioridade:

- Chave API: Global → Projeto → Padrão
- Outras configurações: Projeto → Padrão

Configurações padrão (do código):

- `ignoredPatterns`: `*.exe`, `*.dll`, `*.png`, `*.jpg`, `*.jpeg`, `*.gif`, `*.bin`
- `commitPrefixes`: `chore:`, `refactor:`, `feat:`, `fix:`, `docs:`, `test:`, `etc..`
- `maxCommitLength`: `72`
- `defaultEmoji`: `false`

### Definir sua chave API

```bash
LunaApikey SUA_CHAVE_GEMINI
```

Isso salva a chave no seu `.lunarc` global. Reabra o terminal após definir.

## Uso

Execute Luna dentro de um repositório Git com mudanças staged.

### Comandos e aliases

- `LunaHelp` | `lh` | `-lh`: Mostrar ajuda
- `LunaCommit` | `lc` | `-lc`: Gerar e fazer commits por arquivo
- `LunaApikey <SUA_CHAVE>` | `lkey <SUA_CHAVE>` | `-lkey <SUA_CHAVE>`: Definir chave API
- `LunaConfig` | `config` | `-config` com subcomandos:
  - `init`: Criar `.lunacfg` no diretório atual
  - `show`: Imprimir a configuração mesclada
  - `edit`: Placeholder (não implementado ainda)

Você pode chamá-los como argumentos do executável (ex: `Luna -c`, `Luna lc`, `Luna -lh`) ou diretamente como nomes de comando se seu shell os expuser.

### Fluxo típico

```bash
git add .
Luna -lc          # ou: Luna lc, ou: LunaCommit
```

### Emojis opcionais

```bash
Luna -lc -e       # habilitar emojis nas mensagens
```

## Exemplo de saída

```
Generating commit for file: src/main.go
Committed src/main.go with message:
🚀 feat: add user authentication system

Generating commit for file: README.md
Committed README.md with message:
📝 docs: update installation instructions
```

## Notas

- Luna ignora arquivos binários/imagem comuns
- Se o modelo retornar resposta vazia, fallback é `update <file>`
- Prefixos suportados: `feat:`, `fix:`, `docs:`, `refactor:`, `test:`, `chore:`
- `maxCommitLength` é aplicado (padrão 72)

## Solução de problemas

- Erro: `Set API key using LunaApikey first`
  - Execute `LunaApikey SUA_CHAVE` e reabra o terminal
- Erro executando comandos Git
  - Certifique-se de estar em um repositório Git e que o Git está instalado
- Nenhuma mudança staged
  - Execute `git add .` ou stage arquivos específicos
- Chave API não funcionando
  - Verifique se a chave é válida e tem acesso ao Gemini 2.0 Flash

---

Feito com ❤️ por hax — versão 1.3 (Beta)
