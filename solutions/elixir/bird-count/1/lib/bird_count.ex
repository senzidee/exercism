defmodule BirdCount do
  def today(list) do
    List.first(list)
  end

  def increment_day_count(list) when length(list) == 0, do: [1]
  def increment_day_count(list) do
    [head | tail] = list
    [head + 1 | tail]
  end

  def has_day_without_birds?(list) when length(list) == 0, do: false
  def has_day_without_birds?(list) do
    [head | tail] = list
    head == 0 || has_day_without_birds?(tail)
  end

  def total(list) when length(list) == 0, do: 0
  def total(list) do
    [head | tail] = list
    head + total(tail)
  end

  def busy_days(list) when length(list) == 0, do: 0
  def busy_days(list) do
    [head | tail] = list
    if (head > 4) do
      1 + busy_days(tail)
    else
      busy_days(tail)
    end
  end
end
