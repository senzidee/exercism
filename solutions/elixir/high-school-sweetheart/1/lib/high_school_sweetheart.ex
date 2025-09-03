defmodule HighSchoolSweetheart do
  def first_letter(name) do
    name
    |>String.strip()
    |>String.first()
  end

  def initial(name) do
    first_letter(name)
    |> String.upcase()
    |> Kernel.<>(".")
  end

  def initials(full_name) do
    splitted = full_name
                |> String.split(" ")
    first = splitted
            |> List.first()
            |> initial()
    second = splitted
            |> List.last()
            |> initial()
    first <> " " <> second
  end

  def pair(full_name1, full_name2) do
    """
         ******       ******
       **      **   **      **
     **         ** **         **
    **            *            **
    **                         **
    """
    |> Kernel.<>("**     ")
    |> Kernel.<>(initials(full_name1))
    |> Kernel.<>("  +  ")
    |> Kernel.<>(initials(full_name2))
    |> Kernel.<>("     **\n")
    |> Kernel.<>("""
     **                       **
       **                   **
         **               **
           **           **
             **       **
               **   **
                 ***
                  *
    """)
  end
end
