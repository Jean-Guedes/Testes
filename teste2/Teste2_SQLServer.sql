-- Query para o SQL Server.
-- Criação da tabela temporária de relação entre os dias da semana e código de identificação do cadastro.
-- Será utilizado uma tabela temporária, pois em questão de performance é preferivel utilizar tabelas temporárias do que sub-query.
if object_id('tempdb..#diaCadastro') is not null drop table #diaCadastro
go
select 
	dia.DiaId, hor.CadastroId
into #diaCadastro
from
	DiasSemana as dia
		cross join
        (
			-- Sub-query para recuperação dos códigos de identificação sem repetição.
			select distinct CadastroId from HorariosTrabalho
		) as hor

go

select
	#diaCadastro.CadastroId,
    dia.Descricao as DiaSemana,
    isnull((
		-- Sub-Query que recupera o total de horas trabalhadas em um dia por um cadastro,
        -- Consultando a tabela 'horariostrabalho', utilizando os dados da tabela temporária.
        -- Caso o valor seja nulo, será retornado 0.
		select sum(hor.Horas) from horariostrabalho hor
        where 
			#diaCadastro.CadastroId = hor.CadastroId and
            #diaCadastro.DiaId = hor.DiaId
    ),0) as TotalHoras
from #diaCadastro
inner join diasSemana dia on
	#diaCadastro.DiaId = dia.DiaId
order by
	#diaCadastro.CadastroId,
    #diaCadastro.DiaId

go

-- Remoção da tabela temporária.
drop table #diaCadastro