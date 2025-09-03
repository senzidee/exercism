defmodule Secrets do
  def secret_add(secret) do
    adder = fn param ->
      param + secret
    end
  end

  def secret_subtract(secret) do
    sub = fn param ->
      param - secret
    end
  end

  def secret_multiply(secret) do
    multiplier = fn param ->
      param * secret
    end
  end

  def secret_divide(secret) do
    divider = fn param ->
      floor(param / secret)
    end
  end

  def secret_and(secret) do
    ender = fn param ->
      Bitwise.&&&(param, secret)
    end
  end

  def secret_xor(secret) do
    xorer = fn param ->
      Bitwise.^^^(param, secret)
    end
  end

  def secret_combine(secret_function1, secret_function2) do
    combiner = fn param ->
      secret_function2.(secret_function1.(param))
    end
  end
end
