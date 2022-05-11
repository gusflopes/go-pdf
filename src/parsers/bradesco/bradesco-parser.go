package bradesco

// -----------------------------------------------------------------------------
// Bradesco Parsers
// -----------------------------------------------------------------------------
// ####### Banco > Cliente PF ou PJ > Tipo de Operação > Destinatário (PF ou PJ)
// Obs: Study to see if it's possible to consolidate - Need more example files
//
//
//
// ####### Bradesco > PJ > Mesma Titularidade > PJ
// payee := stringParser(pdfData, "Conta de crédito:", " | CNPJ:")
// // Data => /\d{2}\/\d{2}\/\d{4}/ // First Occurance
// // // Data (Last One) => /\d{2}\/\d{2}/\d{4}(?!.*\d{2}\/\d{2}\/\d{4})/g
// date := regexParser(pdfData, `\d{2}\/\d{2}\/\d{4}(?!.*\d{2}\/\d{2}\/\d{4})`)
// cnpj := regexParser(pdfData, `\d{2}.\d{3}.\d{3}\/\d{4}-\d{2}`)
// amount := regexParser(pdfData, `(?<=Tarifa:R\$ )(.*)(?=Valor)`)

// println("receipt: { date:", date, "amount: ", amount, ", payee: ", payee, ", cnpj: ", cnpj, " }")
// transaction := transaction{Date: date, Amount: amount, Payee: payee, TaxId: cnpj}
