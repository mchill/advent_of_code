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

        file.each_line do |line|
            line = line.split(' ')

            if replacements.key?(line[0])
                replacements[line[0]].push(line[2])
                next
            end

            replacements[line[0]] = [line[2]]
        end

        new_molecules = []
        for atom in 0...molecule.size
            if not replacements.key?(molecule[atom])
                next
            end

            replacements[molecule[atom]].each do |replacement|
                new_molecule = molecule.clone
                new_molecule[atom] = replacement
                new_molecules.push(new_molecule.join)
            end
        end

        puts new_molecules.uniq.count
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
