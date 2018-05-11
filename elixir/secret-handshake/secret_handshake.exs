use Bitwise

defmodule SecretHandshake do
  @doc """
  Determine the actions of a secret handshake based on the binary
  representation of the given `code`.

  If the following bits are set, include the corresponding action in your list
  of commands, in order from lowest to highest.

  1 = wink
  10 = double blink
  100 = close your eyes
  1000 = jump

  10000 = Reverse the order of the operations in the secret handshake
  """
  @spec commands(code :: integer) :: list(String.t())
  def commands(code) when (0b11111 &&& code) == 0, do: []

  def commands(code) when (0b10000 &&& code) > 0 do
    Enum.reverse(commands(code &&& 0b01111))
  end

  def commands(code) when (0b00001 &&& code) > 0 do
    ["wink"] ++ commands(code &&& 0b11110)
  end

  def commands(code) when (0b00010 &&& code) > 0 do
    ["double blink"] ++ commands(code &&& 0b11100)
  end

  def commands(code) when (0b00100 &&& code) > 0 do
    ["close your eyes"] ++ commands(code &&& 0b11000)
  end

  def commands(code) when (0b01000 &&& code) > 0 do
    ["jump"] ++ commands(code &&& 0b10000)
  end

end
