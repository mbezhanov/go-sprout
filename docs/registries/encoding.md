---
description: >-
  The Encoding registry offers methods for encoding and decoding data in
  different formats, allowing for flexible data representation and storage
  within your templates.
---

# Encoding

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`encoding`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/encoding"
```
{% endhint %}

### <mark style="color:purple;">base64Encode</mark>

The function encodes a given string into its Base64 representation, converting the data into a text format suitable for transmission or storage in systems that support Base64 encoding.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Base64Encode(s string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | base64Encode }} // Output: "SGVsbG8gV29ybGQ="
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">base64Decode</mark>

The function decodes a Base64 encoded string back to its original form. If the input string is not valid Base64, it returns an error message.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go"> Base64Decode(s string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "SGVsbG8gV29ybGQ=" | base64Decode }} // Output: "Hello World"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">base32Encode</mark>

The function encodes a given string into its Base32 representation, converting the data into a text format that is compatible with systems supporting Base32 encoding.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Base32Encode(s string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | base32Encode }} // Output: "JBSWY3DPEBLW64TMMQQQ===="
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">base32Decode</mark>

The function decodes a Base32 encoded string back to its original form. If the input is not valid Base32, it returns an error message.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Base32Decode(s string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "JBSWY3DPEBLW64TMMQQQ====" | base32Decode }} // Output: "Hello World"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">fromJson / mustFromJson</mark>

The function converts a JSON string into a corresponding Go data structure, enabling easy manipulation of the JSON data in a Go environment.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">FromJson(v string) any
MustFromJson(v string) (any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ '{"name":"John", "age":30}' | fromJson }} // Output: map[name:John age:30]
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ '{"name":"John", "age":30}' | mustFromJson }} // Output: map[name:John age:30], nil
{{ '{\invalid' | mustFromJson }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toJson / mustToJson</mark>

The function converts a Go data structure into a JSON string, allowing the data to be easily serialized for storage, transmission, or further processing.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToJson(v any) string
MustToJson(v any) (string, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
{{ toJson $d }} // Output: {"key1":"value1","key2":"value2","key3":"value3"}
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
{{ mustToJson $d }} // Output: {"key1":"value1","key2":"value2","key3":"value3"}, nil
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toPrettyJson / mustToPrettyJson</mark>

The function converts a Go data structure into a pretty-printed JSON string, formatting the output with indentation and line breaks for better readability.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToPrettyJson(v any) string
MustToPrettyJson(v any) (string, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
{{ toPrettyJson $d }} // Output: "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\",\n  \"key3\": \"value3\"\n}"
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
{{ mustToPrettyJson $d }} // Output: "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\",\n  \"key3\": \"value3\"\n}"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toRawJson / mustToRawJson</mark>

The function converts a Go data structure into a JSON string without escaping HTML characters, preserving the raw content as it is.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToRawJson(v any) string
MustToRawJson(v any) (string, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "content" "<p>Hello World</p>" }}
{{ toRawJson $d }} // Output: {"content":"<p>Hello World</p>"}
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ $d := dict "content" "<p>Hello World</p>" }}
{{ mustToRawJson $d }} // Output: {"content":"<p>Hello World</p>"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">fromYaml / mustFromYaml</mark>

The function deserializes a YAML string into a Go map, allowing the structured data from YAML to be used and manipulated within a Go program.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">FromYAML(str string) any
MustFromYAML(v string) (any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "name: John Doe\nage: 30" | fromYaml }} // Output: map[name:John Doe age:30]
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ "name: John Doe\nage: 30" | mustFromYAML }} // Output: map[name:John Doe age:30], nil
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toYaml / mustToYaml</mark>

The function serializes a Go data structure into a YAML string, converting the data into a format suitable for YAML representation.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToYAML(v any) string
MustToYAML(v any) (string, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "name" "John Doe" "age" 30 }}
{{ toYaml $d }} // Output: name: John Doe\nage: 30
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ $d := dict "name" "John Doe" "age" 30 }}
{{ $d | mustToYaml }} // Output: name: John Doe\nage: 30
```
{% endtab %}
{% endtabs %}

