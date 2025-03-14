#!/bin/bash

echo "🔧 Configurando ambiente Go..."

# Adiciona GOPATH/bin ao PATH temporariamente
export PATH=$PATH:$(go env GOPATH)/bin
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc # Para ser permanente no ZSH
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc # Para Bash

# Instala as ferramentas necessárias
echo "📦 Instalando ferramentas..."
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/cosmtrek/air@latest

# Verifica se as ferramentas foram instaladas
if command -v swag &> /dev/null && command -v air &> /dev/null
then
    echo "✅ Ambiente configurado com sucesso!"
else
    echo "❌ Algo deu errado na instalação das ferramentas."
fi