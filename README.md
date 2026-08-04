# Suspense Script
### Esse projeto é só de graça
Tava com um problema no Next e fiquei com preguiça de usar IA. Logo pensei:
> _“Por que não fazer um script SÓ pra resolver esse problema que eu vou ter agora e mais nunca?”_

### O problema
Transformei todos os componentes/páginas do meu projeto Next em client components, mas tinha esquecido que só da pra usar o `useSearchParams()` com `<Suspense />`.
Felizmente só usei em 7 arquivos, mas mesmo assim deu preguiça de criar componentes novos só pra envolver em suspense.

E assim surgiu essa ideia **RIDICULA** de fazer esse "script" só pra fazer esse trabalho pra mim :D

> [!NOTE]
> Esse é o meu primeiro projeto com `Go`. Tudo que eu usei aqui foi sabendo o muito básico da linguagem e lendo as documentações — e sem usar IA ;)

### Como rodar
No arquivo `main.go`, altere a seguinte linha para o seu contexto:
```go
appPath := `/home/user/path/src/app/(pages)/`
// ex: /home/user/Documents/suspenseScript/examples/src/app/(pages)/
```
Também é necessário alterar o caminho no arquivo `createComponentFolders.go`
```go
parentDir := fmt.Sprintf("/home/user/path/src/components/%s", fullFolderName)
// ex: /home/user/Documents/suspenseScript/examples/src/components/%s
```

após fazer essas alterações, rode:
```bash
go run main.go
```

## Observações
**O script ainda não está 100%, precisando de MUITAS melhorias e ajustes.**

- `const`s e `import`s ainda precisam ser analisados e alterados, assim como toda a estrutura dos novos componentes e dos arquivos originais(page.tsx).
- Se for usado `[slug]`, é necessário alterar o nome da pasta do componente
- Novos componentes são criados, mas o nome da função **não é alterado**. É necessário fazer essa alteração manualmente.
> [!WARNING]
> A alteração do nome da função deve manter o mesmo nome da pasta, porém com a primeria letra maiúscula:
> ```tsx
> // exemploContainer/index.tsx
> function ExemploContainer() { . . . }
> 
> export default ExemploContainer;
> ```
