# Copyright 2020 Neven Sajko. All rights reserved. See LICENSE for a license.

module FitPlotDat

using DataFrames, Gadfly, Polynomials

export main

const F = BigFloat
const P = Polynomial{F}
const V = Vector{F}
const VI = Vector{Int}

const fpprec = 8192

function readTable()::Tuple{V, V, VI}
	local r::Tuple{V, V, VI} = (F[], F[], Int[])
	while !eof(stdin)
		local latest_seed::Int = parse(Int, readline(), base = 16)
		local distinct_value_count::Int = parse(Int, readline(), base = 10)
		push!(r[1], BigFloat(latest_seed, precision=fpprec))
		push!(r[2], BigFloat(distinct_value_count, precision=fpprec))
		push!(r[3], latest_seed - distinct_value_count + 1)
	end
	r
end

function main()
	setprecision(BigFloat, fpprec)

	local x::V, y::V, d::VI = readTable()

	local line::P = fit(x, y, 1)
	local poly::P = fit(x, y, 47)

	draw(SVG("data/diffs.svg", 50cm, 50cm), plot(
	   Coord.cartesian(xmin=0, ymin=0, ymax=4),
	   layer(DataFrame(latest_seed = x, diff = d), x=:latest_seed, y=:diff)))

	draw(SVG(stdout, 50cm, 50cm), plot(
	   Coord.cartesian(xmin=0, xmax=Int(3e9), ymin=0, ymax=Int(3e9), fixed=true),
	   layer(DataFrame(latest_seed = x, distinct_value_count = y), x=:latest_seed, y=:distinct_value_count),
	   layer([x::F -> line(x), x::F -> poly(x)], 1000, x[length(x)] + 1000, color=[string("linear approximation f(x) = ", line), "approx. with a polynomial of degree 47"])))

	return nothing
end

end

FitPlotDat.main()
