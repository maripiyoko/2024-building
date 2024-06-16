# Ebitengine game jam 2024

## Run game

```
go run .
```

## Build mobile
### Android

```
ebitenmobile bind -target android -javapkg app.chestnuts.mobile.game -o ./mobile/a
ndroid/game.aar ./mobile
```

### iOS

```
ebitenmobile bind -target ios -o ./mobile/ios/Mobile.xcframework ./mobile
```