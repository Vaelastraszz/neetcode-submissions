
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    curr1, curr2 := l1, l2

    var headAdd, currentAdd *ListNode

    // Retenue de l'addition précédente
    carry := 0

    // Continuer tant qu'il reste des chiffres ou une retenue
    for curr1 != nil || curr2 != nil || carry > 0 {

        node := &ListNode{}

        // On commence avec la retenue
        sum := carry

        if curr1 != nil {
            sum += curr1.Val
            curr1 = curr1.Next
        }

        if curr2 != nil {
            sum += curr2.Val
            curr2 = curr2.Next
        }

        // Chiffre à stocker dans le noeud
        node.Val = sum % 10

        // Retenue pour l'itération suivante
        carry = sum / 10

        // Construction de la liste résultat
        if currentAdd == nil {
            currentAdd = node
        } else {
            currentAdd.Next = node
            currentAdd = node
        }

        if headAdd == nil {
            headAdd = node
        }
    }

    return headAdd
}
