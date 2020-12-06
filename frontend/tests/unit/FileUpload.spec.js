import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import FilesUpload from '@/components/FilesUpload'

describe('FilesUpload.vue', () => {
  it('renders props.msg when passed', () => {
    const title = 'new title'
    const wrapper = shallowMount(FilesUpload, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
})
