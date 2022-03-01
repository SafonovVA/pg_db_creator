```
local   all    postgres      trust
local   all    all           md5
```

```postgresql
ALTER USER postgres with password 'postgres';
```