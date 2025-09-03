defmodule RationalNumbers do
  @type rational :: {integer, integer}

  @doc """
  Add two rational numbers
  """
  @spec add(a :: rational, b :: rational) :: rational
  def add(a, b) do
    reduce({elem(a, 0) * elem(b, 1) + elem(b, 0) * elem(a, 1), elem(a, 1) * elem(b, 1)})
  end

  @doc """
  Subtract two rational numbers
  """
  @spec subtract(a :: rational, b :: rational) :: rational
  def subtract(a, b) do
    reduce({elem(a, 0) * elem(b, 1) - elem(b, 0) * elem(a, 1), elem(a, 1) * elem(b, 1)})
  end

  @doc """
  Multiply two rational numbers
  """
  @spec multiply(a :: rational, b :: rational) :: rational
  def multiply(a, b) do
    reduce({elem(a, 0) * elem(b, 0), elem(a, 1) * elem(b, 1)})
  end

  @doc """
  Divide two rational numbers
  """
  @spec divide_by(num :: rational, den :: rational) :: rational
  def divide_by(num, den) do
    den = standardize({elem(den, 1), elem(den, 0)})
    reduce({elem(num, 0) * elem(den, 0), elem(num, 1) * elem(den, 1)})
  end

  @doc """
  Absolute value of a rational number
  """
  @spec abs(a :: rational) :: rational
  def abs(a) do
    reduce({if elem(a, 0) < 0 do elem(a, 0) * -1 else elem(a, 0) end, if elem(a, 1) < 0 do elem(a, 1) * -1 else elem(a, 1) end})
  end

  @doc """
  Exponentiation of a rational number by an integer
  """
  @spec pow_rational(a :: rational, n :: integer) :: rational
  def pow_rational(a, n) do
    if n < 0 do
      reduce({Integer.pow(elem(a, 1), n * -1), Integer.pow(elem(a, 0), n * -1)})
    else
      reduce({Integer.pow(elem(a, 0), n), Integer.pow(elem(a, 1), n)})
    end
  end

  @doc """
  Exponentiation of a real number by a rational number
  """
  @spec pow_real(x :: integer, n :: rational) :: float
  def pow_real(x, n) do
    x ** (elem(n, 0) / elem(n, 1))
  end

  @doc """
  Reduce a rational number to its lowest terms
  """
  @spec reduce(a :: rational) :: rational
  def reduce(a) do
    gcd = Integer.gcd(elem(a, 0), elem(a, 1))
    standardize({trunc(elem(a, 0) / gcd), trunc(elem(a, 1) / gcd)})
  end

  @doc """
  Rewrite rational number in standard format
  """
  @spec standardize(a :: rational) :: rational
  defp standardize(a) do
    {if elem(a, 0) > 0 and elem(a, 1) < 0 do elem(a, 0) * -1 else elem(a, 0) end, if elem(a, 1) < 0 do elem(a, 1) * -1 else elem(a, 1) end}
  end
end
