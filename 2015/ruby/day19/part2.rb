#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        replacements = {}
        molecule = 'CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiR' +
                   'nTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaS' +
                   'iRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCa' +
                   'SiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaS' +
                   'iAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFA' +
                   'rSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFAr' +
                   'PTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiTh' +
                   'SiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRn' +
                   'FYFArCaSiThRnPBPMgAr'
        molecule = molecule.split(/(?=[A-Z])/)

        puts molecule.count - 2 * (molecule.count('Rn') + molecule.count('Y')) - 1
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
