theory Scratch
  imports Main
begin

type_synonym ('key, 'value) Dict = "('key * 'value) set"

fun keys :: "('key, 'value) Dict ⇒ 'key set" where
  "keys d = {k. ∃v. (k, v) ∈ d}"

fun merge :: "('key, 'value) Dict ⇒ ('key, 'value) Dict ⇒ ('key, 'value) Dict" where
  "merge d1 d2 = d1 ∪ {kv ∈ d2. fst kv ∉ keys d1}" 

(* Довести лему, що множина ключів результату злиття є об'єднанням множин ключів аргументів.  *)
lemma merge_keys_disjunction: "keys (merge d1 d2) = keys d2 ∪ keys d1"
  apply auto
  done

(* Довести лему, що якщо множини ключів аргументів не перетинаються, 
   то їх порядок при злитті не важливий - merge A B = merge B A *)
lemma disjoint_map_merge_one:
  assumes  "keys d1 ∩ keys d2 = {}"
  shows "merge d1 d2 = merge d2 d1"
  apply auto
  done
end
