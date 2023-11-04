# exporter_builder
exporter_builder is a project that allows you to create your own prometheus exporter.

## How to use
### install exporter_builder
```bash
go install github.com/k8shuginn/exporter_builder/cmd/builder@latest
```

### create config.yaml
```yaml
name: my_exporter
module: github.com/myname/my_exporter
collectors:
  - sample1
  - sample2
```

### build your exporter
```bash
builder --config ./config.yaml
```
Write code for each collector
