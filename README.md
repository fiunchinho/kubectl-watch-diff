# kubectl watch-diff

A  kubectl plugin that watches the changes in the k8s resource and prints the diff.

## Installation

```bash
go install github.com/alexmt/kubectl-watch-diff@latest
```

## Usage

```bash
kubectl watch-diff cm argocd-cm
--- /var/folders/tz/0snjf0fn797fg_98ttt929tw0000gn/T/kubectl-watch-diff2503959201.yaml	2025-01-25 19:38:20
+++ /var/folders/tz/0snjf0fn797fg_98ttt929tw0000gn/T/kubectl-watch-diff3067301255.yaml	2025-01-25 19:38:20
@@ -13,6 +13,7 @@
         end
     end
     return hs
+  test: hello
 kind: ConfigMap
 metadata:
   annotations:
@@ -24,5 +25,5 @@
     app.kubernetes.io/part-of: argocd
   name: argocd-cm
   namespace: argocd
-  resourceVersion: "303204"
+  resourceVersion: "327879"
   uid: 00a77436-cc07-46d6-9420-3d3037aa85d1
```

[![asciicast](https://asciinema.org/a/699858.svg)](https://asciinema.org/a/699858)