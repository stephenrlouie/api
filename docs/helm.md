# Helm GRPC API

## Chart
    1. Metadata
    2. Template
    3. Chart Deps
    4. Config
    5. Any Files string, bytes

## Config
    1. raw string
    2. map <string, string>

## Metadata
    1. Name string
    2. Home string
    3. Sources string
    4. Description string
    5. Keywords []string
    6. Maintainer
        - Name string
        - Email string
        - URL string
    7. Engine string
    8. Icon string
    9. ApiVersion string
    10. Tags string
    11. AppVersion string
    12. Deprecated bool
    13. TillerVersion string
    14. Annotations map<string, string>
    15. KubeVersion string

## Template
    1. Name string
    2. Data bytes 

## Deps
    1. Name
    2. Repsoitory string
    3. Version string
    4. Alias string
    5. Condition string
    6. Tags: []string


