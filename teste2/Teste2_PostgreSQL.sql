-- Query para o PostgreSQL.
-- Criação da tabela temporária de relação entre os dias da semana e código de identificação do cadastro.
-- Será utilizado uma tabela temporária, pois em questão de performance é preferivel utilizar tabelas temporárias do que sub-query.
drop table if exists dia_cadastro;
create temporary table dia_cadastro as
select 
	dia.DiaId, hor.CadastroId
from
	diassemana as dia
		cross join
        (
			-- Sub-query para recuperação dos códigos de identificação sem repetição.
			select distinct CadastroId from horariostrabalho
		) as hor;

-- Consulta da tabela temporária criada acima, cruzando os dados com as tabela de Horários de trabalho e de Dias da semana. 
select
	dia_cadastro.CadastroId,
    dia.Descricao as DiaSemana,
    Coalesce((
		-- Sub-Query que recupera o total de horas trabalhadas em um dia por um cadastro,
        -- Consultando a tabela 'horariostrabalho', utilizando os dados da tabela temporária.
        -- Caso o valor seja nulo, será retornado 0.
		select sum(hor.Horas) from horariostrabalho hor
        where 
			dia_cadastro.CadastroId = hor.CadastroId and
            dia_cadastro.DiaId = hor.DiaId
    ),0) as TotalHoras
from dia_cadastro
inner join diassemana dia on
	dia_cadastro.DiaId = dia.DiaId
order by
	dia_cadastro.CadastroId,
    dia_cadastro.DiaId;

-- Remoção da tabela temporária.
drop table dia_cadastro;