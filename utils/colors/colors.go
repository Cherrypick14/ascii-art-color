package colors

const (
  reset   = "\033[0m"
  red     = "\033[31m"
  green   = "\033[32m"
  yellow  = "\033[33m"
  blue    = "\033[34m"
  magenta = "\033[35m"
  cyan    = "\033[36m"
)

// GetColor returns the ANSI escape sequence for the provided color name.
func GetColor(color string) string {
  switch color {
  case "red":
    return red
  case "green":
    return green
  case "yellow":
    return yellow
  case "blue":
    return blue
  case "magenta":
    return magenta
  case "cyan":
    return cyan
  default:
    return ""
  }
}
