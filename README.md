# settman
Settings manager for golang from environment.

Features:
 - Allows you to define a default value for settings
 - Allows you to specify mandatory settings
 - Designed to be easy to use and threadsafe
 - Automatically read values from environment and parse into expected type for the setting
 
For clarity, recomend to create a package setting in your program that defines every setting to be used trough code. Then, you can access all settings trough this package.

## Usage:

### Create a new setting:
Arguments to create a new setting are:
  - Name: Specifies the name of the setting, to lookup for this name into environment
  - type: The type of the setting. See at the bottom of this README to check for valid types
  - Default value: If nil, then the setting is mandatory. In case a mandatory setting is not found when parsing, program is going to panic


        OptionalSetting = settman.NewSetting(
            "OPTIONAL_SETTING_NAME",
            settman.Uint8,
            uint8(3),
          )

        MandatorySetting = settman.NewSetting(
            "MANDATORY_SETTING_NAME",
            settman.Bool,
            nil,
          )
 

### Get value for a setting:
Before getting a value, you must parse it. Just do:

    OptionalSetting.Parse()
    MandatorySetting.Parse()
    
Recomend to parse this setting at the beggining of your code, but you can call parse as many times you want, to get the value from environment
Once the setting is parsed, you can use it like this:

    val := OptionalSetting.Get().(uint8)
    val2 := MandatorySetting.Get().(bool)
    
**You can allways asume that Get is going to return a valid value if you have already parsed this setting. You can safely convert to the expected type without the need to check if it's valid or not**

### Full example:

    package settings

    import (
      "github.com/ackFacu/settman"
      "github.com/joho/godotenv"
    )

    var (
      LogLevel = settman.NewSetting(
        "LOG_LEVEL",
        settman.Uint8,
        uint8(1),
      )

      Env = settman.NewSetting(
        "ENV",
        settman.String,
        nil,
      )

      Port = settman.NewSetting(
        "PORT",
        settman.Uint32,
        uint32(8080),
      )
    )

    func init() {
      // Load .env file if exists
      _ = godotenv.Load()

      LogLevel.Parse()
      Env.Parse()
      Port.Parse()
    }
    
    
Then, you can access the settings like this:

    package main

    import (
      "fmt"
      "github.com/ackFacu/gotemplate/settings"
    )

    func main() {

      if settings.LogLevel.Get().(uint8) < uint8(2) {
        fmt.Printf("hello. Running on environment %s, on port %d", settings.Env.Get().(string), settings.Port.Get().(uint32))
      }
    }
