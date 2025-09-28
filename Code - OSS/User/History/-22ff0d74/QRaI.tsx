'use client'
import Image from 'next/image'
import { useEffect, useState, useRef } from 'react'

import logo from '../../../public/logos/logo.svg'
import search from '../../../public//logos/search.svg'
import bag from '../../../public//logos/shopping_bag.svg'
import account from '../../../public/logos/account_circle.svg'

import ModalAccount from '../components/modalAccount'

export default function Header(){
    const [openAcc, setOpenAcc] = useState(false);
    const handleModalAcc = () => setOpenAcc(!openAcc);

    {/* CLICK OUTSIDE BUTTON AREA */}

    const modalRef = useRef<HTMLDivElement>(null);
    useEffect(() => {
        const handleClickOut = (event: MouseEvent) => {
            if ( openAcc &&
                ref.current && !ref.current.contains(event.target)
            ){
                setOpenAcc(false);
            }
        };

        document.addEventListener('mousedown', handleClickOut);

        return () => {
            document.removeEventListener('mousedown', handleClickOut);
        };

    }, [openAcc, setOpenAcc]);


    return (
    <div className="bg-light">
        <div className="flex justify-center w-full h-10 bg-dark-gray mb-2">
            <p className="text-light text-sm font-bold font-noto content-center ">UMA REVOLUÇÃO NO SEU GUARDA ROUPA</p>
        </div>

        <div className="flex mx-10 h-15 justify-between">
            <div className="content-center">
                <a href= "#">
                <Image src={logo} alt="Urban Soul" />
                </a>
            </div>

            <div className="grid grid-flow-col gap-4 content-center font-noto text-sm text-dark-gray">
                <a href="#" className="hover:font-bold hover:underline underline-offset-4 decoration-red decoration-2" >Masculino</a> 
                <a href="#" className="hover:font-bold hover:underline underline-offset-4 decoration-red decoration-2" >Feminino</a>
                <a href="#" className="hover:font-bold hover:underline underline-offset-4 decoration-red decoration-2" >Kids</a>
            </div>

            <div className="content-center">
            <div className="flex">
                <div className="flex bg-light-gray w-40 h-5 rounded-2xl hover:border-1">
                    <Image className="ml-2" src={search} alt=""/>
                    <input type="text" className="ml-2 text-gray font-bold font-noto text-xs content-center outline-none "
                            name="filter" placeholder="Buscar" size={12} />
                </div>

                <div className="flex ml-4 gap-2">
                    <a href="#">
                        <Image src={bag} alt=""/>
                    </a>
                    <button className="cursor-pointer" onClick={handleModalAcc}>
                        <Image src={account} alt=""/>
                    </button>
                </div>
            </div>
            </div>
        </div>

        <div className="flex justify-end">
            <ModalAccount ref={modalRef} open={openAcc}></ModalAccount>
        </div>
    </div>
    )
}