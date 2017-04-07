package main

import (
  "fmt"
  "math"
)
// Returns the damage and the Shield HP

func shieldCalculation (damageTaken int, shieldHP int, shieldArmor int) (int, int) {
  if shieldHP <= 0 {
    shieldHP = 0
    return damageTaken, shieldHP
  } else if shieldArmor >= damageTaken {
    if shieldHP < damageTaken {
      damageTaken -= shieldHP
      shieldHP = 0
      return damageTaken, shieldHP
    } else {
      shieldHP -= damageTaken
      damageTaken = 0
      return damageTaken, shieldHP
    }
  } else {
    shieldHP -= shieldArmor
    damageTaken -= shieldArmor
    return damageTaken, shieldHP
  }
}

// Returns true or false

func criticalCalculator(criticalChance int) bool {
  criticalGenerator := randomNumber(100)
  if criticalGenerator <= criticalChance {
    return true
  } else {
    return false
  }
}

// Returns the damage maded

func damageCalculator (damageTaken int, enemyArmor int, enemyLife int) int {
  if damageTaken <= enemyArmor {
    return 1
  } else {
    return (damageTaken - enemyArmor)
  }
}

// Returns enemy's life, enemy's shield HP, damage maded and if was critic

func attack(damage int, enemyArmor int, criticalChance int, enemyLife int, enemyShieldArmor int, enemyShieldHP int) (int, int, int, bool) {
  critic := criticalCalculator(criticalChance)
  if critic {
    damage *= 2
  }
  damage, enemyShieldHP = shieldCalculation(damage, enemyShieldHP, enemyShieldArmor)
  damage = damageCalculator(damage, enemyArmor, enemyLife)
  enemyLife -= damage
  return enemyLife, enemyShieldHP, damage, critic
}

func fleeCalculation(HPPorcent float64) float64 {
  HPPorcent = math.Floor(HPPorcent)
  fleePorcent := math.Floor(HPPorcent * 75 / 100)
  return fleePorcent
}

func flee(HPPorcent float64) {
  fleePorcent := fleeCalculation(HPPorcent)
  fleeRandom := randomNumber(100)
  if fleeRandom <= int(fleePorcent) {
    fmt.Println("¡Has escapado!")
    //goToTown()
  } else {
    fmt.Println("No pudiste escapar")
    //nextTurn()
  }
}
