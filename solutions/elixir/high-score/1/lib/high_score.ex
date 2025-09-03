defmodule HighScore do
  def new(), do: %{}

  def add_player(scores, name, score \\ 0), do: Map.put(scores, name, score)

  def remove_player(scores, name), do: Map.delete(scores, name)

  def reset_score(scores, name), do: add_player(remove_player(scores, name), name)

  def update_score(scores, name, score) do
    if (scores[name] == nil) do
        add_player(remove_player(scores, name), name, score)
      else
        add_player(remove_player(scores, name), name, score + scores[name])
    end
  end

  def get_players(scores) do
    Map.keys(scores)
  end
end
