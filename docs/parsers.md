## Parsers


### BRADESCO - TED Mesma Titularidade - Beneficiário PJ - Conta PJ
```go
Valor Líquido => "Tarifa:R$ ", "Valor",
Beneficiário => "Conta de crédito:", " | CNPJ: "
CNPJ => " | CNPJ: ", "Empresa:"
Data => /\d{2}\/\d{2}\/\d{4}/ // First Occurance
Data => /\d{2}\/\d{2}/\d{4}(?!.*\d{2}\/\d{2}\/\d{4})/g // Last Ocurrance


```