---
description: >-
  The Std registry provides a set of standard functions for common tasks,
  included by default, making it easy to perform basic operations without
  additional setup.
---

# Std

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`std`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/std"
```
{% endhint %}

### <mark style="color:purple;">hello</mark>

The function returns a simple greeting string, "Hello!" It serves as a basic test function to verify that the system is working correctly.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Hello() string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>{{ hello }} // Output: Hello!
</strong></code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">default</mark>

The function returns the first non-empty value from a provided list of arguments. If the list is empty or the first value is empty, it returns a specified default value. If you're looking to find the first non-empty value from a list of multiple options, the [`Coalesce` ](std.md#coalesce)function is a better choice.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Default(defaultValue any, given ...any) any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ nil | default "default" }} // Output: "default"
{{ "" | default "default" }}  // Output: "default"
{{ "first" | default "default" }} // Output: "first"
{{ "first" | default "default" "second" }} // Output: "second"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">empty</mark>

The function checks if the provided value is empty, returning `true` if it is considered empty based on its type. This function is useful for determining whether a value is present or absent across different data types.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Empty(given any) bool
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ nil | empty }} // Output: true
{{ "" | empty }} // Output: true
{{ 0 | empty }} // Output: true
{{ false | empty }} // Output: true
{{ struct{}{} | empty }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">all</mark>

The function checks if all values in the provided variadic slice are non-empty. It returns `true` only if every value is considered non-empty according to the criteria used by the [`Empty` ](std.md#empty)method.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">All(values ...any) bool
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 1, "hello", true | all }} // Output: true
{{ 1, "", true | all }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">any</mark>

The function checks if any of the provided values are non-empty. It returns `true` if at least one value is considered non-empty.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Any(values ...any) bool
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "", 0, false | any }} // Output: false
{{ "", 0, "text" | any }} // Output: true
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">coalesce</mark>

The function returns the first non-empty value from the provided list. If all values are empty, it returns `nil`.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Coalesce(values ...any) any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ nil, "", "first", "second" | coalesce }} // Output: "first"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ternary</mark>

The function mimics the ternary conditional operator found in many programming languages. It returns `trueValue` if the `condition` is true; otherwise, it returns `falseValue`.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Ternary(trueValue any, falseValue any, condition bool) any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ true | ternary "yes", "no" }} // Output: "yes"
{{ false | ternary "yes", "no" }} // Output: "no"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">cat</mark>

The function concatenates a series of values into a single string, converting each value to its string representation and separating them with spaces. Nil values are skipped, and no trailing spaces are added.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Cat(values ...any) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello", nil, 123, true | cat }} // Output: "Hello 123 true"
```
{% endtab %}
{% endtabs %}
