defmodule KitchenCalculator do
  def get_volume(volume_pair) do
    elem(volume_pair, 1)
  end

  def to_milliliter(volume_pair) when elem(volume_pair, 0) == :cup, do: {:milliliter, elem(volume_pair, 1) * 240}
  def to_milliliter(volume_pair) when elem(volume_pair, 0) == :fluid_ounce, do: {:milliliter, elem(volume_pair, 1) * 30}
  def to_milliliter(volume_pair) when elem(volume_pair, 0) == :teaspoon, do: {:milliliter, elem(volume_pair, 1) * 5}
  def to_milliliter(volume_pair) when elem(volume_pair, 0) == :tablespoon, do: {:milliliter, elem(volume_pair, 1) * 15}
  def to_milliliter(volume_pair) when elem(volume_pair, 0) == :milliliter, do: volume_pair

  def from_milliliter(volume_pair, unit) when unit == :cup, do: {:cup, elem(volume_pair,1) / 240}
  def from_milliliter(volume_pair, unit) when unit == :fluid_ounce, do: {:fluid_ounce, elem(volume_pair,1) / 30}
  def from_milliliter(volume_pair, unit) when unit == :teaspoon, do: {:teaspoon, elem(volume_pair,1) / 5}
  def from_milliliter(volume_pair, unit) when unit == :tablespoon, do: {:tablespoon, elem(volume_pair,1) / 15}
  def from_milliliter(volume_pair, unit) when unit == :milliliter, do: volume_pair

  def convert(volume_pair, unit), do: from_milliliter(to_milliliter(volume_pair), unit)
end
