<template>
  <div class="p-4">
    <button @click="signPermit" class="bg-blue-600 text-white px-4 py-2 rounded">
      Sign USDC Permit
    </button>
    <div v-if="signature" class="mt-4 break-all text-sm">
      ✅ Signature: {{ signature }}
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { ethers } from 'ethers'

export default {
  setup() {
    const signature = ref(null)

    const signPermit = async () => {
      if (!window.ethereum) return alert('Please install MetaMask')

      // 使用 BrowserProvider 和 MetaMask 连接
      const provider = new ethers.BrowserProvider(window.ethereum)
      const signer = await provider.getSigner()
      const address = await signer.getAddress()

      // 从后端获取 TypedData
      const res = await fetch('http://localhost:8080/api/permit-typed-data?owner=' + address)
      const typedData = await res.json()

      // 使用 MetaMask 发起签名请求
      const signatureHex = await window.ethereum.request({
        method: 'eth_signTypedData_v4',
        params: [address, JSON.stringify(typedData)],
      })

      // 保存签名
      signature.value = signatureHex
      console.log('✅ Signed permit:', signatureHex)

      // 提交签名到后端
      await fetch('http://localhost:8080/api/submit-permit', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ signature: signatureHex, typedData }),
      })
    }

    return {
      signature,
      signPermit,
    }
  }
}
</script>
